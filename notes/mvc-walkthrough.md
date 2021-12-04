## Walking Through a Web Request with MVC

MVC may be easier to understand if we take a normal web request and examine how each part of the request would be handled. To do this, we will walk through a web request where a user is attempting to update their contact information.

<!-- TODO: MVC Figure 1-2 -->

**1. A user submits an update to their contact information**

This would typically happen when the user goes to their account settings and updates a few fields there. For example, they might update their name and email address. This web request is then sent to our server, where the router is the first code to take action.


**2. The router routes to the UserController**

When the router gets the web request, it realizes that this is a request to update a User's contact information based on the URL and [HTTP method](http://www.restapitutorial.com/lessons/httpmethods.html), so it forwards the request along to the `UserController` to handle the request.

<!-- TODO: MVC Figure 3-4 -->

**3. The UserController uses the UserStore to update the user's contact info**

While handling the request, the `UserController` will need to make a change to data stored in our database, but our controllers don't interact directly with the database. Instead the controller will use the `UserStore` provided by the `models` package to update the contact information, which allows us to isolate all of our database specific code so that our controllers don't need to worry about it.


**4. The UserStore returns the updated data**

After updating the user in the database, the UserStore will then return the updated user object to the controller so that the controller can use this in the rest of its code.

<!-- TODO: MVC Figure 5-6 -->

**5. The UserController uses the ShowUser view to generate HTML**

Once the controller gets the user back, it can then use the `ShowUser` view generate an HTML response showing the update. Notice that our controller is basically just calling code from the `views` and `models` packages. It rarely does work on its own, and instead acts more as a director.


**6. The ShowUser view writes and HTML response to the user**

After everything is said and done, our `ShowUser` view will have rendered the HTML page that shows a user and the end user will now see their user account with the contact information updated. Mission success!
