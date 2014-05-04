part of G0;

abstract class Api{

  /**
   * Loads images from [page] async and returns result as [Map];
   */
  Future<Map> getImages({int offset: 0, int count: 20});
}
