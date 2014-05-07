part of G0;

/**
 * Shows single image and handles sizing
 */
class Detail{

  // DOM Elements
  Element _element;
  Element _cover;
  Element _spinner;
  Element _imageContainer;
  Element _footer;
  Element _user;
  Element _channel;
  Element _date;
  Element _source;

  // Image meta
  String id;
  String imageUrl;
  String source;
  String user;
  String channel;
  int date;

  int _windowWidth;
  ImageElement _loadedImage;

  DateFormat _dateFormat = new DateFormat(G0.DATE_FORMAT);

  Detail(){
    _getElements();
    _onResize();
    _eventBindings();
  }

  /**
   * Search document for neccessary DOM elements
   */
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

  void _eventBindings(){
    _cover.onClick.listen((_) => _hideDetail());
    window.onResize.listen((_) => _onResize());
    window.onKeyUp.listen(_handleKeys);
  }

  /**
   * Saves window width on resize and starts image resizing
   */
  void _onResize(){
    _windowWidth = window.innerWidth;
    _setImageSize();
  }

  /**
   * Takes a [LIElement] from .image-list retrievs meta from it and displays
   * singe image
   */
  void show(LIElement target){

    _retrieveImageMeta(target);

    assert(id != null);
    assert(imageUrl != null);

    _loadedImage = new ImageElement(src: imageUrl);
    _loadedImage.classes.add('detail-image');
    _loadedImage.onLoad.listen((Event evt) => _showImage(evt.target));

    DateTime imgDate = new DateTime.fromMillisecondsSinceEpoch(date);

    _user.innerHtml = user;
    _channel.innerHtml = channel;
    _date.innerHtml = '${_dateFormat.format(imgDate)}';
    _source.innerHtml = source;
    _source.setAttribute('href', source);

    window.history.pushState(
        null,
        imageUrl,
        window.location.pathname + '?offset=$id'
    );

    _showCover();
    _showDetail();
  }

  void _retrieveImageMeta(LIElement target){
    id = target.dataset['id'];
    imageUrl = target.dataset['image'];
    source = target.dataset['source'];
    user = target.dataset['user'];
    channel = target.dataset['channel'];
    date = int.parse(target.dataset['date']) * 1000 ;
  }


  /**
   * Find [LIElement] by offset and delegates to [show]
   * This is used for offset from queryString
   */
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
    _imageContainer.innerHtml = '';
    _hideCover();

    window.history.pushState(
        null,
        imageUrl,
        window.location.pathname
    );
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
    _imageContainer.append(_loadedImage);
    _spinner.classes.remove('show');
  }

  /**
   * Calculate and apply optimal image size
   */
  void _setImageSize(){
    if(_loadedImage == null){
      return;
    }
    int origWidth = int.parse(_loadedImage.dataset['width']);
    int origHeight = int.parse(_loadedImage.dataset['height']);

    int width = origWidth > _windowWidth ? _windowWidth : origWidth;
    double ratio = width / origWidth;
    int height = (origHeight * ratio).ceil();
    int left = width >= _windowWidth ? 0 : ((_windowWidth - width) / 2).ceil();

    _element..style.width = '${width}px'
            ..style.height = '${height}px'
            ..style.left = '${left}px'
            ..style.top = '120px';
  }

  void _handleKeys(KeyboardEvent evt){
    switch(evt.keyCode){
      case 27:
        _hideDetail();
        break;
    }
  }
}
