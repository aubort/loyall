
/**
 * Handles the contact form submit
 */
$("#contact-form").submit(function(e) {
    var url = "/contactus/";
    console.log($("#contact-form").serialize());
    
    $.ajax({
          type: "POST",
          url: url,
          data: $("#contact-form").serialize(), // serializes the form's elements.
          success: function(data)
            {
              alert(data); // show response from the php script.
            }
         });
    
    e.preventDefault(); // avoid to execute the actual submit of the form.
});
