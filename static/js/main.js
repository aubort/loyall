
/**
 * Handles the contact form submit
 */
$("#contact-form").submit(function(e) {
    var url = "/contactus/";
    var formData = $("#contact-form").serializeObject();
    
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


$.fn.serializeObject = function()
{
    var o = {};
    var a = this.serializeArray();
    $.each(a, function() {
        if (o[this.name] !== undefined) {
            if (!o[this.name].push) {
                o[this.name] = [o[this.name]];
            }
            o[this.name].push(this.value || '');
        } else {
            o[this.name] = this.value || '';
        }
    });
    return o;
};