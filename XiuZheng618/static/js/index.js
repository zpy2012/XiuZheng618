window.onload = function() {
	//强制手机竖屏访问
	(function rotate() {
		var orientation = window.orientation;
		if(orientation == 90 || orientation == -90) {
			document.body.style.display = 'none';
			alert("请使用竖屏访问！");
		}
		window.onorientationchange = function() {
			document.body.style.display = "block";
			rotate();
		};
	})()
};
