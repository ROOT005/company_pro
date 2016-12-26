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
				console.log(data)
				if (data == "false"){
					alert("请输入正确的验证码");
					$(".captcha-img").trigger('click');
				}else{
					$(".verify").addClass('true');
				}
			}
		});
	});
	$(".phonenum").change(function(){
		var phonenum = $(".phonenum").val();
		$.ajax({
			type:"get",
			url:"/submit/phonenum?phonenum="+phonenum,
			async:true,
			success:function(data){
				console.log(data);
				if (data == "false") {
					alert("请输入正确的电话");
				}else{
					$(".phonenum").addClass('true');
				}
			}
		});
	});
	$('.submit').click(function(){
		var result =  $('.verify').hasClass('true') && $('.phonenum').hasClass('true');
		if (!result) {
			alert("请按要求填写资料");
		}
		return result;
	});
});
