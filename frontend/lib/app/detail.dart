part of G0;

class Detail{

  Element _element;
  Element _cover;
  Element _spinner;
  Element _imageContainer;

  Element _footer;
  Element _user;
  Element _channel;
  Element _date;
  Element _source;

  int _windowWidth;
  ImageElement loadedImage;

  DateFormat dateFormat = new DateFormat(G0.DATE_FORMAT);

  Detail(){
    _getElements();
    _onResize();

    if(_cover != null){
      _cover.onClick.listen((_) => _hideDetail());
    }

    window.onResize.listen((_) => _onResize());
  }

  void _getElements(){
    _element = querySelector('#container .detail');
    _cover = querySelector('#container .cover');
    _spinner = _element.querySelector('.spinner');
    _footer = _element.querySelector('.footer');
    _user = _footer.querySelector('.user');
    _channel = _footer.querySelector('.channel');
    _date = _footer.querySelector('.date');
    _source = _footer.querySelector('.source');
    _imageContainer = _element.querySelector('.image-container');
  }

  void _onResize(){
    _windowWidth = window.innerWidth;
    _setImageSize();
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
    int date = int.parse(target.dataset['date']) * 1000 ;

    assert(id != null);
    assert(imageUrl != null);

    loadedImage = new ImageElement(src: imageUrl);
    loadedImage.classes.add('detail-image');
    loadedImage.onLoad.listen((Event evt) => _showImage(evt.target));

    DateTime imgDate = new DateTime.fromMillisecondsSinceEpoch(date);

    _user.innerHtml = user;
    _channel.innerHtml = channel;
    _date.innerHtml = '${dateFormat.format(imgDate)}';
    _source.innerHtml = source;
    _source.setAttribute('href', source);

    window.history.pushState(
        null,
        'imageUrl',
        window.location.pathname + '?offset=$id'
    );

    _showCover();
    _showDetail();
  }

  void showByOffset(String offset){
    LIElement target = querySelector('.image-list li[data-id="$offset"]');
    if(target != null){
      show(target);
    }
  }

  void _showDetail(){
    _spinner.classes.add('show');
    _element.classes.add('show');
    _footer.classes.add('show');

  }

  void _hideDetail(){
    _spinner.classes.remove('show');
    _element.classes.remove('show');
    _footer.classes.remove('show');
    _hideCover();
  }

  void _showCover(){
    _cover.classes.add('show');
  }

  void _hideCover(){
    _cover.classes.remove('show');
  }

  void _showImage(ImageElement img){
    img.dataset['width'] = img.width.toString();
    img.dataset['height'] = img.height.toString();

    _setImageSize();
    _imageContainer..innerHtml = ''
                   ..append(loadedImage);

    _spinner.classes.remove('show');
  }

  void _setImageSize(){
    if(loadedImage == null){
      return;
    }
    int origWidth = int.parse(loadedImage.dataset['width']);
    int origHeight = int.parse(loadedImage.dataset['height']);

    int width = origWidth > _windowWidth ? _windowWidth : origWidth;
    double ratio = width / origWidth;
    int height = (origHeight * ratio).ceil();
    int left = width >= _windowWidth ? 0 : ((_windowWidth - width) / 2).ceil();

    _element..style.width = '${width}px'
                   ..style.height = '${height}px'
                   ..style.left = '${left}px'
                   ..style.top = '120px';
  }
}
