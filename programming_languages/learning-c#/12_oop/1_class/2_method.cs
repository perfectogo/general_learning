using System;
using System.Security.Cryptography.X509Certificates;

class Program
{   
    static void Main()
    {
        // User user = new User();

        // user.Username = "ali";
        // Console.WriteLine(user.Username);

       // string passoword = user.GetPassword();

        var newUser = new User("ali", "1001");

        Console.WriteLine(newUser.GetToken());
    }

}

public class User 
{
    public string Username {set; get;}
    private string password = "0000";
    private Guid token;

    public User () 
    {

    }

    public User (string username, string passoword) 
    {
        Username = username;
        this.password = passoword;
        token = Guid.NewGuid();

    }

    public string GetPassword() {
        return this.password;
    }

    public Guid GetToken() {
        return this.token;
    }
}
