import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.regex.Pattern;
import java.util.regex.Matcher;
import java.util.List;
import java.util.ArrayList;

public class Day3 {
    public static void main(String[] args) throws Exception {
        String fileContent = readFile();
        part1(fileContent);
        // part2(fileContent);
    }

    private static String readFile() throws Exception {
        Path p = Paths.get(System.getProperty("user.home"), "Desktop", "Projects", "aoc24", "day3", "input.txt");
        return Files.readAllLines(p)
            .stream()
            .collect(StringBuilder::new, StringBuilder::append, StringBuilder::append)
            .toString();
    }

    private static void part1(String content) {
        String re = "(?<=mul\\()(\\d{1,3},\\d{1,3})(?=\\))";
        Pattern pattern = Pattern.compile(re);
        Matcher matcher = pattern.matcher(content);
        List<String> matches = new ArrayList<>();
        while (matcher.find()) {
            matches.add(matcher.group());
        }
        int sum = matches.stream()
            .map(s -> s.split(","))
            .mapToInt(pair -> Integer.parseInt(pair[0]) * Integer.parseInt(pair[1]))
            .sum();
        System.out.println(sum);
    }

    // private static void part2(String content) {
    //     String re = "";
    //     Pattern pattern = Pattern.compile(re);
    //     Matcher matcher = pattern.matcher(content);
        


    // }
}