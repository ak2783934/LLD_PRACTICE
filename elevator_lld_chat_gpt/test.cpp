vector<vector<string>> getProductSuggestion(vector<string> products, string search) {
    // Sort products lexicographically
    sort(products.begin(), products.end());

    vector<vector<string>> result;
    string prefix = "";

    for (char ch : search) {
        prefix += ch;
        vector<string> suggestions;

        // Find first index where product >= prefix
        auto it = lower_bound(products.begin(), products.end(), prefix);

        // Collect up to 3 matching products
        for (int i = 0; i < 3 && it + i != products.end(); i++) {
            if ((it + i)->compare(0, prefix.size(), prefix) == 0) {
                suggestions.push_back(*(it + i));
            } else {
                break;
            }
        }

        result.push_back(suggestions);
    }

    return result;
}