class Solution {
    /**
     * @param {character[][]} board
     * @return {boolean}
     */
 check(row) {
  const res = new Set();

  for (const x of row) {
    if (x === ".") continue;
    if (res.has(x)) return false;
    res.add(x);
  }

  return true;
}

 isValidSudoku(board) {
  // rows
  for (const row of board) {
    if (!this.check(row)) {
      console.log("a");
      return false;
    }
  }

  // cols
  const columns = Array.from({ length: 9 }).map((_, i) =>
    board.map((v) => v[i]),
  );

  for (const col of columns) {
    if (!this.check(col)) {
      console.log("b");
      return false;
    }
  }

  // boxes

  const bStart = [
    [0, 0],
    [0, 3],
    [0, 6],
    [3, 0],
    [3, 3],
    [3, 6],
    [6, 0],
    [6, 3],
    [6, 6],
  ];

  const boxes = bStart.map(([a, b]) =>
    [
      [a, b],
      [a, b + 1],
      [a, b + 2],
      [a + 1, b],
      [a + 1, b + 1],
      [a + 1, b + 2],
      [a + 2, b],
      [a + 2, b + 1],
      [a + 2, b + 2],
    ].map(([a, b]) => board[a][b]),
  );

  for (const b of boxes) {
    if (!this.check(b)) {
      console.log("c");
      return false;
    }
  }

  return true;
}

}
