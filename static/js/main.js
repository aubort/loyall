
/**
 * Handles the contact form submit
 */
$("#contact-form").submit(function(e) {
    var url = "/contactus/";
    var formData = $("#contact-form").serialize();
    
    $.ajax({
          type: "POST",
          url: url,
          data: formData, // serializes the form's elements.
          success: function(data)
            {
            //   alert(data); // show response from the php script.
            }
         });
    
    e.preventDefault(); // avoid to execute the actual submit of the form.
});

var verifyRecaptchaCallback = function(response) {
  $("#submit-button").prop('disabled', false);
};