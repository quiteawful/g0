part of G0;

abstract class Api{

  /**
   * Loads images from [page] async and returns result as [Map];
   */
  Future<Map> getImages({int page});
}
