jQuery.noConflict();
jQuery(function($){
	$(".verify").change(function(){
		var id = $('.verify').val();
		var captcha_id = $('.cap_id input').val();
		$.ajax({
			type: "get",
			url:"/verify?captcha_id="+ captcha_id+"&id="+id,
			async:true,
			success:function(data){
				message = data;
				if (message == "false"){
					alert("请输入正确的验证码");
					$(".captcha-img").trigger('click');
				}else{
					console.log(message);
				}
			}
		});
	});
});