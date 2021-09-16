import javax.annotation.Resource;
import java.lang.reflect.Field;
import java.util.Date;
import java.util.HashMap;
import java.util.Map;
// 示例实体类
class Entity {

    // 尝试从SpringIOC容器中获取名为name的Bean
    @Resource
    private String name;
    // 尝试从SpringIOC容器中获取名为name的Bean
    @Resource
    private Date date;
    private int age;

    public int getAge() {
        return age;
    }

    public String getName() {
        return name;
    }

    public Date getDate() {
        return date;
    }

    @Override
    public String toString() {
        return "Entity{" +
                "name='" + name + '\'' +
                ", date=" + date +
                ", age=" + age +
                '}';
    }
}

class FakeSpring {
    private static final Map<String, Object> container = new HashMap<>();
    public static void addBean(String name, Object bean) {
        container.put(name, bean);
    }
    public static Object getBean(String name) {
        return container.get(name);
    }

    public static <T> T getBean(String name, Class<T> c) {
        return (T)getBean(name);
    }

    /**
     * 对传入的对象的字段填充SpringIOC容器中对应的Bean
     * @param obj
     */
    public static void pad(Object obj) throws IllegalAccessException {
        Field[] fields = obj.getClass().getDeclaredFields();
        for (Field field : fields) {
            // 若字段中存在@Resource注解，就把对应的Bean弄进去
            if (field.getAnnotation(Resource.class) != null) {
                String name = field.getName();
                field.setAccessible(true);
                Object value = container.get(name);
                field.set(obj, value);
            }
        }
    }
    public static void main(String[] args) throws IllegalAccessException {

        // <bean />@Bean @Service @Component @Controller @RestController @Configuration
        /**
         * 相当于
         * <bean name="name" class="java.lang.String">
         *     <constructor>
         *         <param>xiaotao</param>
         *     </constructor>
         * </bean>
         * 或者是相当于
         * @Bean
         * public String name() {
         *     return "xiaotao";
         * }
         */
        FakeSpring.addBean("name", "xiaotao");
        /**
         * 相当于
         * <bean name="name" class="java.util.Date"></bean>
         * 或者是相当于
         * @Bean
         * public Date date() {
         *     return new Date();
         * }
         */
        FakeSpring.addBean("date", new Date());

        // 创建空对象
        // @Component
        Entity e = new Entity();
        // 字段注入填充
        FakeSpring.pad(e);
        FakeSpring.addBean("entity", e);
        System.out.println(e);
    }
}
