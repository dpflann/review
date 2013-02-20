/**
 * Class implementing the class game: Tic-Tac-Toe.
 *  i = column, j = row
 * j   X |   | O
 *    -----------
 *     O | X |
 *    -----------
 *     O |   | X
 */

import java.util.Scanner;
import java.util.regex.Pattern;
import java.util.InputMismatchException;

public class TicTacToe {

  protected static final int X = 1, O = -1;
  protected static final int EMPTY = 0;
  protected int board[][] = new int[3][3];
  protected int player;

  public TicTacToe() {
    clearBoard();
  }

  public void clearBoard() {
    for (int i=0; i < 3; i++) {
      for (int j=0; j < 3; j++) {
        board[i][j] = EMPTY;
      }
    }
    player = O;
  }

  public boolean putMark(int i, int j) throws IllegalArgumentException {
    if ((i < 0) || (i > 2) || (j < 0) || (j > 2)) {
      throw new IllegalArgumentException("Invalid board position");
    }

    if (board[i][j] == EMPTY) {
      board[i][j] = player;
      player = -1 * player;
      return true;
    }

    else {
      return false;
    }

  }

  public boolean isWin(int mark) {
    //short circuit this, small board
    return (
        equalsThreeMarks(board[0][0] + board[1][0] + board[2][0], mark) || //check horizontal
        equalsThreeMarks(board[0][1] + board[1][1] + board[2][1], mark) ||
        equalsThreeMarks(board[0][2] + board[1][2] + board[2][2], mark) ||
        equalsThreeMarks(board[0][0] + board[0][1] + board[0][2], mark) || //check vertical
        equalsThreeMarks(board[1][0] + board[1][1] + board[1][2], mark) ||
        equalsThreeMarks(board[2][0] + board[2][1] + board[2][2], mark) ||
        equalsThreeMarks(board[0][0] + board[1][1] + board[2][2], mark) || //check diagonal 00 to 22
        equalsThreeMarks(board[0][2] + board[1][1] + board[2][0], mark)); //check diagonal 02 to 20
  }

  public boolean equalsThreeMarks(int sum, int mark) {
    return sum == 3 * mark;
  }

  public int winner() {
    if (isWin(X))
      return X;
    else if (isWin(O))
      return O;
    else
      return 0;
  }

  public String toString() {
    String s = "";
    for (int i=0; i < 3; i++) {
      for (int j=0; j < 3; j++) {
        switch (board[i][j]) {
          case X: s += "  X  "; break;
          case O: s += "  O  "; break;
          case EMPTY: s += "(" + i + "," + j + ")"; break;
        }
        if (j < 2)
          s += "|";
      }
      if (i < 2)
        s += "\n-----------------\n";
    }
    return s;
  }

  public String getPlayer() {
    if (player == X)
      return "X";
    else
      return "O";
  }

  public boolean noWinner() {
    return  winner() == 0;
  }

  public boolean boardIsOpen() {
    int zeros = 0;

    for (int i=0; i < 3; i++) {
      for (int j=0; j < 3; j++) {
        if (board[i][j] == EMPTY)
          ++zeros;
      }
    }

    return zeros != 0;
  }

  public int[] columnAndRow(String location) {
    int[] cAndR = new int[2];

    if (location.indexOf(',') > -1) {
      String[] iAndJ = location.split(",");
      cAndR[0] = Integer.parseInt(iAndJ[0]);
      cAndR[1] = Integer.parseInt(iAndJ[1]);
    }
    else {
      cAndR[0] = Integer.parseInt(location.substring(0,1));
      cAndR[1] = Integer.parseInt(location.substring(1));
    }
    return cAndR;
  }

  public void printWinner() {
    int winner = winner();
    String announcement = "";
    switch (winner) {
      case X: announcement = "X is the winner!"; break;
      case O: announcement = "O is the winner!"; break;
      case 0: announcement = "Stalemate, cat's game, no winner, you get the point"; break;
    }
    System.out.println(announcement);
  }

  public static void main(String[] args) {
    Scanner scanner = new Scanner(System.in);
    Pattern location = Pattern.compile("[0-2]{1},?[0-2]{1}");
    TicTacToe ticTacToe = new TicTacToe();

    System.out.println("Get your TicTacToe on!");
    System.out.println("Two players are needed, find a friend quickly.\n");
    System.out.println("Player as " + ticTacToe.getPlayer() + " goes first\n");
    System.out.println("The current board is blank:\n" + ticTacToe +"\n");
    System.out.println("Place your piece in the position desired by typing in the column followed by a comma then the row\ne.g. the middle (1,1) --> input: 1,1");

    while (ticTacToe.noWinner() && ticTacToe.boardIsOpen()) {
      System.out.println(ticTacToe.getPlayer() + "'s turn\n");
      String s = scanner.next(location);
      int[] cAndR = ticTacToe.columnAndRow(s);

      if (!ticTacToe.putMark(cAndR[0], cAndR[1]))
        System.out.println("That space is not empty, try again.");
      System.out.println("\n" + ticTacToe + "\n");
    }//end while

    ticTacToe.printWinner();
  }

} // end TicTacToe
