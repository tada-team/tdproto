abstract class IResponse<T> {
  String get time;

  bool get ok;

  T get result;

  String get error;

  Map<String, dynamic> get details;
}
