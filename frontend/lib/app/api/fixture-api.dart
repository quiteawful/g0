part of G0;

/**
 * Fixture version of [Api]. Use this for testing
 */
class FixtureApi implements Api {

  Map result;

  FixtureApi(int perPage){
    result = {
      'page' : 1,
      'image-src' : 'assets/images/',
      'thumb-src' : 'assets/images/',
      'images': []
    };

    for(var i = 0; i < perPage; i++){
      Map item = {
        'image': 'test1.jpg',
        'thumb': 'testThumb1.jpg',
      };

      result['images'].add(item);
    }
  }

  Future<Map> getImages({int page}){
    Completer compl = new Completer();
    Future f = new Future.delayed(
        new Duration(seconds: 1),
        () => compl.complete(result)
    );
    return compl.future;
  }
}
