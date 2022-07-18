# Example Java client

```java
String superintendentAddress = "localhost";
int superintendentPort = 556644;

try (Socket socket = new Socket(superintendentAddress, superintendentPort)) {

    // Send a example "event"
    PrintWriter writer = new PrintWriter(socket.getOutputStream(), true);
    writer.println("{\"example\":\"this is an example event\"}");

    // Listen to "events" and print them into the log
    Scanner scan = new Scanner(new InputStreamReader(socket.getInputStream()));
    while(scan.hasNextLine()){
         String line = scan.nextLine();
         System.out.println(line);
    }

} catch (Exception ex) {
    System.out.println("error: " + ex.getMessage());
}
```