# Chapter 19 - 메서드

1. 리시버가 있으면 메서드, 없으면 일반 함수입니다.
2. 리시버는 메서드를 호출하는 주체로써 메서드는 리시버를 통해서만 호출할 수 있습니다.
    - 따라서 메서드는 리시버에 속한 기능을 표현합니다.
    - 모든 로컬 타입은 리시버가 될 수 있습니다.
3. 메서드를 통해서 데이터와 기능이 묶임으로써 객체라는 개념이 생겼고, 프로그래밍 패러다임은 순서도 위주의 절차 중심에서 객체 사이의 관계 중심으로 변화했습니다.
4. **포인터 메서드는 인스턴스 중심으로 메서드에서 호출자 인스턴스에 접근하여 값을 변경**할 수 있습니다.
5. 값 타입 메서드 호출 시 값이 모두 복사됩니다. 인스턴스가 아닌 값 중심의 메서드를 만들 때 사용합니다.
    - 호출자 인스턴스에 접근할 수 없고 복사되는 양에 따라서 성능상 문제가 될 수 있습니다.
