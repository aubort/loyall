
/**
 * Handles the contact form submit
 */
$("#contact-form").submit(function(e) {
  e.preventDefault(); // avoid to execute the actual submit of the form.
    var url = "/contactus/";
    var formData = $("#contact-form").serialize();
    if (validateForm(formData)){
      return;
    }
    
    $("#submit-button").prop('disabled', true);
    $("#submit-button").prop('value', 'Sending...');
    
    $.ajax({
          type: "POST",
          url: url,
          data: formData, // serializes the form's elements.
          success: function(data)
            {
              window.location.href = "/confirmation/";
              //console.log("OK!");
            //   alert(data); // show response from the php script.
            },
          error: function(data){
            //console.log("ERROR");
          }
        });
    
    
});

var verifyRecaptchaCallback = function(response) {
  $("#submit-button").prop('disabled', false);
};

var validateForm = function(formData){
  
  var isDirty = false;
  var inputs = document.querySelectorAll(".mdl-textfield");
    for (var i = 0, l = inputs.length; i < l; i++) {
      if(inputs[i].MaterialTextfield.input_.value == ""){
        inputs[i].classList.add("is-invalid", "is-dirty");
        isDirty = true;
      }
    }
  
  return isDirty;
}