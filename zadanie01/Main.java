

public class Main {
    public static void main(String[] args) {
        System.out.println("Hello, World!");

        try {
            Class.forName("org.sqlite.JDBC");
            System.out.println("SQLite JDBC is installed");
        } catch (ClassNotFoundException e) {
            System.out.println("SQLite JDBC needs to be installed");
            e.printStackTrace();
        }
    }
}
