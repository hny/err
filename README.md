### Quick use:
```golang

import Err "github.com/hny/err"

func foo() Err.Err {
    err := ……
    if err != nil {
        return Err.New().Err(
            "……",err,
        )
    }
    
    return nil
}

```

### if you wangt to return the default error:
```golang

import Err "github.com/hny/err"

func foo() error {
    err := ……
    if err != nil {
        return Err.New().Err(
            "……",err,
        )
    }

    return nil
}

```