import sys
from stats import get_num_words, get_num_letters, chars_dict_to_sorted_list

def get_book_text(filepath: str) -> str:
    """
    Function to retrieve the text of a book.
    """
    with open(filepath, 'r') as file:
        return file.read()

def print_report(filepath: str, word_count: int, letter_count: list):
    """
    Function to print the report of the book analysis.
    """
    print("============ BOOKBOT ============")
    print(f"Analyzing book found at {filepath}...")
    print("----------- Word Count ----------")
    print(f"Found {word_count} total words")
    print("--------- Character Count -------")
    for letter, count in letter_count:
        if letter.isalpha():
            print(f"{letter}: {count}")
    print("============= END ===============")

def main():
    """
    Main function to run the book bot.
    """
    if len(sys.argv) < 2:
        print("Usage: python3 main.py <path_to_book>")
        sys.exit(1)
    filepath = sys.argv[1]
    book_text = get_book_text(filepath)
    word_count = get_num_words(book_text)
    letter_count = chars_dict_to_sorted_list(get_num_letters(book_text))
    print_report(filepath, word_count, letter_count)

if __name__ == "__main__":
    main()