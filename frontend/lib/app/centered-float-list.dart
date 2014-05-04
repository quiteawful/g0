part of G0;

/**
 * Keeps [_imageList] centered. Needs at least one item inside of [_imageList]
 * Todo: replace this class with proper CSS solution
 */
class CenteredFloatList{

  Element _imageList;
  int _containerWidth;
  int _itemWidth;

  bool _isInitialized = false;
  bool get isInitialized => _isInitialized;

  CenteredFloatList(this._imageList){
    init();
    window.onResize.listen(_center);
  }

  /**
   * Checks for items in [_imageList] and initializes based on first item width.
   * Does nothing when no item is found.
   */
  void init(){
    if(_imageList == null){
      return;
    }

    List<Element> elements = _imageList.querySelectorAll('li.image-list-item');
    if(elements.length != 0){
      _itemWidth = elements.first.offsetWidth;
      _isInitialized = true;
      _center(null);
    }
  }

  /**
   * adds margin to [_imageList] to center it inside of its parent if [this]
   * is initialized.
   */
  void _center(Event evt){
    if(!_isInitialized){
      return;
    }
    _containerWidth = _imageList.parent.offsetWidth;
    int margin = _containerWidth % _itemWidth;
    _imageList.style.margin = '0 ${margin/2}px';
  }
}
