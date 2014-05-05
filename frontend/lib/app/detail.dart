part of G0;

class Detail{

  Element _element;
  Element _cover;

  Detail(){
    _element = querySelector('#container .detail');
    _cover = querySelector('#container .cover');

    if(_cover != null){
      _cover.onClick.listen((_) => _hideDetail());
    }
  }

  void show(Element target){
    if(_element == null || _cover == null){
      return;
    }

    String id = target.dataset['id'];
    String imageUrl = target.dataset['image'];
    String source = target.dataset['source'];
    String user = target.dataset['user'];
    String channel = target.dataset['channel'];
    String date = target.dataset['date'];

    assert(id != null);
    assert(imageUrl != null);

    _showCover();
  }

  void _hideDetail(){
    _hideCover();
  }

  void _showCover(){
    _cover.classes.add('show');
  }

  void _hideCover(){
    _cover.classes.remove('show');
  }
}
