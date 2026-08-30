def get_num_words(book_text: str) -> int:
    """
    Function to count the number of words in a book.
    """
    return len(book_text.split())

def get_num_letters(book_text: str) -> dict[str, int]:
    """
    Function to count the number of letters in a book.
    """
    letter_count = {}
    for letter in book_text:
        letter = letter.lower()
        if letter in letter_count:
            letter_count[letter] += 1
        else:
            letter_count[letter] = 1
    return letter_count

def sort_on(letter_count: tuple[str, int]) -> int:
    """
    Function to sort the letter count dictionary.
    """
    return letter_count[1]

def chars_dict_to_sorted_list(letter_count: dict[str, int]) -> list[tuple[str, int]]:
    """
    Function to convert the letter count dictionary to a sorted list.
    """
    letter_count = [(letter, count) for letter, count in letter_count.items()]
    return sorted(letter_count, key=sort_on, reverse=True)