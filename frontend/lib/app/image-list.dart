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

  String host = window.location.hostname;

  /**
   * id of last loaded image
   */
  String lastId = '';

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
      lastId = data['id'].toString();
      Future delayed = new Future.delayed(
          new Duration(milliseconds: delay),
          () => item.classes.add('loaded')
      );
      delay += 50;
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
    AnchorElement a = new AnchorElement(href: '${_imageSrc}${data['image']}');
    ImageElement img = new ImageElement(src: '${_imageSrc}${data['thumb']}');
    li.classes.add('image-list-item');
    li.dataset['data-id'] = data['id'].toString();
    a.append(img);
    li.append(a);
    return li;
  }

  /**
   * Updates [_pageWidth], [_pageHeight] and [perPage]. This is called after
   * resize events
   */
  void _getPerPage(){
    _pageWidth = window.innerWidth;
    _pageHeight = window.innerHeight - _imageList.offsetTop;
    perRow = (_pageWidth / _imageWidth).floor();

    int newPerPage = (_pageWidth / _imageWidth).floor()
                   * (_pageHeight / _imageHeight).ceil();

    if(newPerPage != perPage){
      print('change to $newPerPage images per page');
      perPage = newPerPage;
    }
  }
}