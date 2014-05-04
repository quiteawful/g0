part of G0;

/**
 * Creates and displays images inside of [_imageList].
 */
class ImageList {

  Element _imageList;
  Element _loadingElement;

  /**
   * Keps track of loaded page
   */
  int currentPage = 0;

  String _imageSrc;
  String _thumbSrc;

  /**
   * Stores all loaded items
   */
  List<LIElement> items = new List<LIElement>();

  ImageList(this._imageList){
    if(_imageList != null){
      _loadingElement = _imageList.querySelector('.loading');
    }
  }

  /**
   * Creates new items based on [api] result [data]
   * an appends them to [_imageList]
   */
  void showImages(Map data){
    _imageSrc = data['image-src'];
    _thumbSrc = data['thumb-src'];
    currentPage = data['page'];

    int delay = 0;
    data['images'].forEach((image){
      Element item = createItem(image);
      _imageList.append(item);
      items.add(item);
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
    a.append(img);
    li.append(a);
    return li;
  }
}
