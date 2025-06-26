// @target: esnext

class C {
  protected constructor() {}
}

class B<T = any> extends C {}

class A extends B {
  f() {
    new A();
  }
}