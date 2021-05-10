# Chapter 10 - switch문

```go
switch 비교값 {
case 값1:
    문장                    // 비교값과 값1이 같을 때 수행
case 값2:
    문장                    // 비교값과 값2가 같을 때 수행
default:                    // default 구문은 생략 가능
    문장                    // 만족하는 case가 없을 때 수행
}
```

```go
switch {                    // 비교값을 생략하면 비교값은 true
case 조건문1:
    문장                    // 조건문1이 true면 실행
case 조건문2:
    문장                    // 조건문2가 true면 실행
default:                    // default 구문은 생략 가능
    문장                    // 모든 case의 조건문이 false면 실행
}
```

```go
switch 초기문; 비교값 {     // 초기문이 먼저 실행되고 비교값을 case들과 비교
case 값1:
    문장                    // 비교값과 값1이 같으면 실행
case 값2:
    문장                    // 비교값과 값2가 같으면 실행
default:                    // default 구문은 생략 가능
    문장                    // 만족하는 case가 없을 때 수행
}
```
