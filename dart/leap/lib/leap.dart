class Leap {
  bool leapYear(name) {
    return name % 400 == 0 || (name % 100 != 0 && name % 4 == 0);
  }
}
