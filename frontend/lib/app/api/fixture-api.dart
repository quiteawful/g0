part of G0;

/**
 * Fixture version of [Api]. Use this for testing
 */
class FixtureApi implements Api {

  Map result;
  int currentId = 0;

  FixtureApi(){
    result = {
      'image-src' : 'assets/images/',
      'thumb-src' : 'assets/images/',
      'images': []
    };
  }

  Future<Map> getImages({String offset: '', int count: 20}){
    print('load $count images from offset: $offset');
    result['images'].clear();
    for(var i = 0; i < count; i++){
      currentId++;
      Map item = {
        'id': currentId,
        'image': 'test1.jpg',
        'thumb': 'test-thumb1.png',
      };

      result['images'].add(item);
    }

    Completer compl = new Completer();
    Future f = new Future.delayed(
        new Duration(seconds: 1),
        () => compl.complete(result)
    );
    return compl.future;
  }
}
