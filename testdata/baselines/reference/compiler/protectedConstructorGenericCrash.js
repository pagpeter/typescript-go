//// [tests/cases/compiler/protectedConstructorGenericCrash.ts] ////

//// [protectedConstructorGenericCrash.ts]
class C {
  protected constructor() {}
}

class B<T = any> extends C {}

class A extends B {
  f() {
    new A();
  }
}

//// [protectedConstructorGenericCrash.js]
class C {
    constructor() { }
}
class B extends C {
}
class A extends B {
    f() {
        new A();
    }
}
