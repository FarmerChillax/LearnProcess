use pyo3::prelude::*;

/// Formats the sum of two numbers as string.
#[pyfunction]
fn sum_as_string(a: usize, b: usize) -> PyResult<String> {
    Ok((a + b).to_string())
}

fn m(x: usize) -> usize {
    return x * x * x * x * x;
}

///
#[pyfunction]
fn test_func() -> PyResult<String> {
    for a in 0..10 {
        for b in 0..10 {
            for c in 0..10 {
                for d in 0..10 {
                    for e in 0..10 {
                        for f in 0..10 {
                            if m(a) + m(b) + m(c) + m(d) + m(e) + m(f)
                                == a * 10_0000 + b * 1_0000 + c * 1000 + d * 100 + e * 10 + f
                            {
                                println!("{}{}{}{}{}{}", a, b, c, d, e, f);
                            }
                        }
                    }
                }
            }
        }
    }
    return Ok(String::new());
}

/// A Python module implemented in Rust.
#[pymodule]
fn string_sum(py: Python, m: &PyModule) -> PyResult<()> {
    m.add_function(wrap_pyfunction!(sum_as_string, m)?)?;
    m.add_function(wrap_pyfunction!(test_func, m)?)?;

    Ok(())
}
