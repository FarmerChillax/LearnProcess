class PromiseCore {
    static PENDING = "pending";
    static FULFILLED = "fulfilled";
    static REJECTED = "rejected";
    // init instance
    constructor(executor) {
      this.status = PromiseCore.PENDING;
      this.value = null;
      this.callbacks = []
      try {
        executor(this.resolve.bind(this), this.reject.bind(this));
      } catch (error) {
        this.reject(error);
      }
    }

    resolve(value) {
      if (this.status == PromiseCore.PENDING) {
        this.status = PromiseCore.FULFILLED;
        this.value = value;
        setTimeout(() => {
            this.callbacks.map(callback => {
                callback.onFulfilled(value)
            })
        });
      }
    }

    reject(value) {
      if (this.status == PromiseCore.PENDING) {
        this.status = PromiseCore.REJECTED;
        this.value = value;
        setTimeout(() => {
            this.callbacks.map(callback => {
                callback.onRejected(value)
            })
        });
      }
    }

    then(onFulfilled, onRejected) {
        if (typeof onFulfilled != "function") {
          onFulfilled = value => value;
        }
        if (typeof onRejected != "function") {
          onRejected = value => value;
        }

        return new PromiseCore((resolve, reject) => {
            if (this.status == PromiseCore.PENDING) {
                // push to callbacks
                this.callbacks.push({
                    // 对象的 field
                    onFulfilled: (value) => {
                        try { 
                            // method arg
                            let result = onFulfilled(value)
                            resolve(result)
                        } catch (error) {
                            onRejected(error)
                        }
                    }, 
                    onRejected: value => {
                        try {
                            let result = onRejected(value)
                            resolve(result)
                        } catch (error) {
                            onRejected(error)
                        }
                    }
                })
            }
    
            if (this.status == PromiseCore.FULFILLED) {
              setTimeout(() => {
                try {
                    let result = onFulfilled(this.value);
                    resolve(result)
                  } catch (error) {
                    onRejected(error);
                  }
              });
            }
            if (this.status == PromiseCore.REJECTED) {
              setTimeout(() => {
                try {
                    let result = onRejected(this.value);
                    resolve(result)
                  } catch (error) {
                    onRejected(error);
                  } 
              });
            }           
        })
      }
  }