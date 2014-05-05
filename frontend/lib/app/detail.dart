part of G0;

class Detail{

  Element _element;
  Element _cover;
  Element _spinner;
  Element _imageContainer;

  int _windowWidth;

  Detail(){
    _element = querySelector('#container .detail');
    _cover = querySelector('#container .cover');
    _spinner = _element.querySelector('.spinner');
    _imageContainer = _element.querySelector('.image-container');
    _setWindowSize();

    if(_cover != null){
      _cover.onClick.listen((_) => _hideDetail());
    }

    window.onResize.listen((_) => _setWindowSize());

  }

  void _setWindowSize(){
    _windowWidth = window.innerWidth;
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

    ImageElement img = new ImageElement(src: imageUrl);
    img.classes.add('detail-image');
    img.onLoad.listen((Event evt) => _showImage(evt.target));

    _showCover();
    _showDetail();
  }

  void _showDetail(){
    _spinner.classes.add('show');
    _element.classes.add('show');
  }

  void _hideDetail(){
    _spinner.classes.remove('show');
    _element.classes.remove('show');
    _hideCover();
  }

  void _showCover(){
    _cover.classes.add('show');
  }

  void _hideCover(){
    _cover.classes.remove('show');
  }

  void _showImage(ImageElement img){
    _spinner.classes.remove('show');

    _imageContainer..innerHtml = ''
                   ..append(img);

  }
}
