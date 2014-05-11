part of G0;

/**
 * Creates and displays images inside of [_imageList].
 */
class ImageList {

  Element _imageList;
  Element _loadingElement;
  int _imageWidth;
  int _imageHeight;
  int _pageWidth;
  int _pageHeight;

  Detail detail = new Detail();

  String host = window.location.hostname;

  /**
   * id of last loaded image
   */
  String currentOffset = '';

  /**
   * active element
   */
  LIElement active;

  /**
   * Number of images which fit in one screen
   */
  int perPage = 0;
  int perRow = 0;

  String _imageSrc;
  String _thumbSrc;

  /**
   * Stores all loaded items
   */
  List<LIElement> items = new List<LIElement>();

  ImageList(this._imageList, this._imageWidth, this._imageHeight){
    _getPerPage();
    if(_imageList != null){
      _loadingElement = _imageList.querySelector('.loading');
    }
    window.onResize.listen((_) => _getPerPage());
    window.onKeyDown.listen((KeyboardEvent evt){
      switch(evt.keyCode){
        case 37: //Left
        case 65: //A
          prev();
          break;
        case 39: //Right
        case 68: //D
          next();
          break;
      }
    });

  }

  /**
   * Creates new items based on [api] result [data]
   * an appends them to [_imageList]
   */
  void showImages(Map data){
    _imageSrc = data['image-src'];
    _thumbSrc = data['thumb-src'];

    int delay = 0;
    data['images'].forEach((data){
      Element item = createItem(data);
      _imageList.append(item);
      items.add(item);
      currentOffset = data['id'].toString();

      //Click event for detail view
      item.onClick.listen(_onImageClick);

      Future delayed = new Future.delayed(
          new Duration(milliseconds: delay),
          () => item.classes.add('loaded')
      );
      delay += 30;
    });
  }

  /**
   * Makes loading spinner visible
   */
  void showLoading(){
    if(_loadingElement != null){
       _loadingElement.classes.add('loading');
     }
  }

  /**
   * Makes loading spinner invisible
   */
  void hideLoading(){
    if(_loadingElement != null){
      _loadingElement.classes.remove('loading');
    }
  }

  /**
   * Takes [data] and returns HTML representation [Element]
   */
  Element createItem(Map data){
    LIElement li = new LIElement();
    AnchorElement a = new AnchorElement(href: '${_imageSrc}${data['img']}');
    ImageElement img = new ImageElement(src: '${_imageSrc}${data['thumb']}');
    li.classes.add('image-list-item');
    li.dataset['id'] = data['id'].toString();
    li.dataset['user'] = data['user'].toString();
    li.dataset['channel'] = data['channel'].toString();
    li.dataset['date'] = data['date'].toString();
    li.dataset['source'] = data['source'].toString();
    li.dataset['image'] = '${_imageSrc}${data['img'].toString()}';

    a.append(img);
    li.append(a);
    return li;
  }

  /**
   * Updates [_pageWidth], [_pageHeight] and [perPage]. This is called after
   * resize events
   */
  void _getPerPage(){

    // bad hack. We need hight offset to get enaugh images to force scrollbar
    int heightOffset = 20;

    _pageWidth = window.innerWidth;
    _pageHeight = window.innerHeight - _imageList.offsetTop + heightOffset;
    perRow = (_pageWidth / _imageWidth).floor();

    int newPerPage = (_pageWidth / _imageWidth).floor()
                   * (_pageHeight / _imageHeight).ceil();

    if(newPerPage != perPage){
      print('change to $newPerPage images per page');
      perPage = newPerPage;
    }
  }

  void _onImageClick(Event evt){
    evt.preventDefault();
    Element target = evt.target;
    if(target is !LIElement){
      target = target.parentNode;
    }
    _setActive(target);
    detail.show(target);
  }

  void _setActive(LIElement target){
    items.forEach((LIElement item){
      item.classes.remove('active');
    });
    target.classes.add('active');
    active = target;
  }

  void next(){
    int index;
    if(active == null){
      index = 1;
    } else {
      index = items.indexOf(active) + 1;
    }
    if(index < items.length){
      LIElement target = items[index];
      _setActive(target);
      detail.show(target);
    }
  }

  void prev(){
    int index;
    if(active == null){
      index = 0;
    } else {
      index = items.indexOf(active) - 1;
    }
    if(index >= 0){
      LIElement target = items[index];
      _setActive(target);
      detail.show(target);
    }
  }
}
