/**
 * Implementing basic and more advanced pattern matching algorithms
 **/

public class PatternMatching
{

  /* Brute Force, O(n^2) */
  public static boolean BruteForcePatternMatch(String text, String pattern)
  {

    if (text.length() < pattern.length())
      return false;

    char[] textChars = text.toCharArray();
    char[] patternChars = pattern.toCharArray();

    for (int i = 0; i <= (textChars.length - patternChars.length); i++)
    {
      int j = 0;
      for (; j < patternChars.length; j++)
        if (patternChars[j] != textChars[i + j])
          break;
      if (j == patternChars.length)
        return true;
    }
    return false;
  }

  /* Boyer-Moore Algorithm */
  public static boolean BoyerMoorePatternMatch(String text, String pattern)
  {
    int[] alphabetLastIndexes = lastIndexFunction(pattern);
    int textLength = text.length();
    int patternLength = pattern.length();
    int textIndex = patternLength - 1;
    int patternIndex = patternLength - 1;

    if (patternIndex > textLength - 1)
      return false;

    do
    {
      if (pattern.charAt(patternIndex) == text.charAt(textIndex))
      {
        if (patternIndex == 0)
          return true;
        else //continue character comparison from back to front
        {
          textIndex--;
          patternIndex--;
        }
      }
      else //character mismatch, move the window on the text with regard to the last index occurrence of current letter in text
      {
        textIndex = textIndex + patternLength - Math.min(patternIndex, 1 + alphabetLastIndexes[text.charAt(textIndex)]);
        patternIndex = patternLength - 1;
      }

    } while (textIndex <= textLength - 1); //can reach end of text for final comparison

    return false;
  }

  /* build table for all possible characters in the alphabet for the strings
   * value is -1 if character is not in the pattern, otherwise, it is the
   * index of the last occurence of the character in the pattern
   */
  private static int[] lastIndexFunction(String pattern)
  {
    int[] alphabetLastIndexes = new int[128];
    for (int i = 0; i < 128; i++)
      alphabetLastIndexes[i] = -1;
    for (int i = 0; i < pattern.length(); i++)
      alphabetLastIndexes[pattern.charAt(i)] = i;

    return alphabetLastIndexes;
  }


  public static void main(String[] args)
  {
    String a = "ababcd";
    String b = "abc";
    System.out.println("Brute Force Pattern Matching");
    System.out.println("\"" + b + "\" in \"" + a + "\"? " + (BruteForcePatternMatch(a, b) ? "Yes." : "No."));
    b = "def";
    System.out.println("\"" + b + "\" in \"" + a + "\"? " + (BruteForcePatternMatch(a, b) ? "Yes." : "No."));

    System.out.println("Boyer-Moore Pattern Matching");
    b = "abc";
    System.out.println("\"" + b + "\" in \"" + a + "\"? " + (BoyerMoorePatternMatch(a, b) ? "Yes." : "No."));
    b = "def";
    System.out.println("\"" + b + "\" in \"" + a + "\"? " + (BoyerMoorePatternMatch(a, b) ? "Yes." : "No."));
  }
}// end class
