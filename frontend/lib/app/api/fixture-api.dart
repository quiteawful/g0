part of G0;

/**
 * Fixture version of [Api]. Use this for testing
 */
class FixtureApi implements Api {

  Map result;
  int currentId = 0;
  Random rand = new Random();

  List fixtureItems = [
    {
      'id':      '',
      'img':     'test1.jpg',
      'thumb':   'test-thumb1.png',
      'source':  'http://www.lucrorn.com',
      'user':    'coke',
      'channel': '#winebottle',
      'date':    '1399332098'
    },
    {
      'id':      '',
      'img':     'test2.jpg',
      'thumb':   'test-thumb2.jpg',
      'source':  'http://www.nuthing.com',
      'user':    'kern',
      'channel': '#winebottle',
      'date':    '1399392098'
    },
    {
      'id':      '',
      'img':     'test3.jpg',
      'thumb':   'test-thumb3.jpg',
      'source':  'http://zziellos.com',
      'user':    'ziellos',
      'channel': '#winebottle',
      'date':    '1394392098'
    }
  ];

  FixtureApi(){
    result = {
      'image-src' : 'res/images/',
      'thumb-src' : 'res/images/',
      'images': []
    };
  }

  Future<Map> getImages({String offset: '', int count: 20}){
    print('load $count images from offset: $offset');
    result['images'].clear();
    for(var i = 0; i < count; i++){
      currentId++;
      Map item = fixtureItems[rand.nextInt(3)];
      item['id'] = currentId;

      result['images'].add(item);
    }

    Completer compl = new Completer();
    Future f = new Future.delayed(
        new Duration(milliseconds: 300),
        () => compl.complete(result)
    );
    return compl.future;
  }
}
