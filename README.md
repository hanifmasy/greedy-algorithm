# greedy-algorithm

This program defines a `Food` struct to represent each food item with its name, carbohydrate content, and cost per 100 grams. It then calculates the cost per gram of carbohydrate for each food item.

After that, it iterates through all possible combinations of rice, corn, and potato quantities (in multiples of 100 grams) such that the total carbohydrate content equals 400 grams. For each combination, it calculates the total cost and keeps track of the combination with the minimum cost.

Finally, it outputs the combination of foods with the minimum cost and the total cost.

This method applies a brute force approach to solve a combinatorial optimization problem. Here's how it can be categorized in terms of programming method:

1. **Brute Force**: The program exhaustively tries all possible combinations of quantities for each food item to find the combination that minimizes the total cost while meeting the constraint of providing 400 grams of carbohydrates. This approach works well for small problem sizes, as it guarantees finding the optimal solution without sophisticated algorithms.

2. **Combinatorial Optimization**: The problem involves optimizing a combination of discrete variables (quantities of different foods) subject to a constraint (total carbohydrate content). The program systematically explores the solution space and evaluates the cost of each combination to identify the one that satisfies the constraint and has the lowest cost.

3. **Greedy Algorithm**: Although not explicitly implemented as a greedy algorithm, the approach of calculating the cost per gram of carbohydrate for each food item and then selecting the combination with the lowest cost can be seen as a form of greedy strategy. However, the brute force search over all combinations makes it less efficient than typical greedy algorithms.

Overall, this method demonstrates a straightforward way to tackle optimization problems by systematically exploring the solution space. While effective for small instances, more efficient algorithms (such as dynamic programming or heuristic methods) might be necessary for larger problem sizes.
