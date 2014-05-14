part of G0;

/**
 * Fires event when scrollPosition is at end of window
 */
class InfinteLoad {

  Element _target;
  int _loadOffset;
  int _loadDelay;

  int _windowHeight;
  int _scrollY = 0;
  int _targetOffset;
  int _targetHeight;

  Stopwatch _stopwatch = new Stopwatch()..start();

  /**
   * Event gehts only fired when this is true.
   * We need this to prevent recursive loading
   */
  bool isActivated = true;

  StreamController _onFire = new StreamController();
  Stream get onFire => _onFire.stream;

  /**
   * [loadOffset] defines scroll offset. This is used to fire event before
   * we reache en of window
   */
  InfinteLoad(this._target, {int loadOffset: 150, int loadDelay: 300}){
      if(_target == null){
        return;
      }
      this._loadOffset = loadOffset;
      this._loadDelay = loadDelay;

      _windowHeight = window.innerHeight;
      updateTargetHeight();

      window.onScroll.listen(onScroll);
      window.onResize.listen(onResize);
  }

  /**
   * Handles scroll events and triggers [onFire] stream.
   */
  void onScroll(Event evt){
    _scrollY = window.scrollY;
    if(_stopwatch.elapsedMilliseconds < _loadDelay){
      return;
    }

    _stopwatch.reset();
    if(_scrollY + _windowHeight >= _targetHeight + _targetOffset - _loadOffset){
      if(isActivated){
        _onFire.add(true);
        isActivated = false;
      }
    }
  }

  /**
   * Handles resize events to update window and [_target] height and offset.
   */
  void onResize(Event evt){
    _windowHeight = window.innerHeight;
    updateTargetHeight();
  }

  /**
   * Updates height and offset of target.
   */
  void updateTargetHeight(){
    _targetHeight = _target.offsetHeight;
    _targetOffset = _target.offsetTop;
  }

  /**
   * Needs to be called after every finished load.
   */
  void activate(){
    isActivated = true;
  }
}
