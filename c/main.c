bool isPalindrome(char* s) {
    int end_index = strlen(s);
    int j = end_index-1;
    for (int i = 0; i < end_index && j > 0;)
    {
        while (!isalnum(s[i]) && i < end_index)
        {
            i++;
        }
        while (!isalnum(s[j]) && j > 0)
        {
            j--;

        }
        if (i >= j)
        {
            return true;
        }
        if (tolower(s[i]) != tolower(s[j]))
        {
            return false;
        }
        i++;
        j--;
    }
    return true;
}
