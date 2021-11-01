abstract class IWebsocketEvent<T> {
  String get event;

  T get params;

  String get confirmId;
}
