#include <iostream>

//СТРУКТУРА ЭЛЕМЕНТА СПИСКА
struct ListNode
{
    int val;
    ListNode* next;

    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode* next) : val(x), next(next) {}
};

// АЛГОРИТМ
ListNode* ReverseList(ListNode* head)
{
    if (head->next == nullptr) return head;
 
    ListNode* end = head;
    ListNode* current = head->next;
    ListNode* temp = head->next;

    while (temp != nullptr)
    {
        temp = temp->next;
        current->next = head;
        head = current;
        current = temp;
    }

    end->next = nullptr;
    return head;
}

// ВЫВЕСТИ СПИСОК НА КОНСОЛЬ
void showList(ListNode* list)
{
    while (list != nullptr)
    {
        std::cout << list->val << " ";
        list = list->next;
    }
}

// ТЕСТ
int main()
{
    std::cout << "Test...\n";

    ListNode* prev = new ListNode(10);
    for (size_t i = 9; i != 0; i--)
    {
        ListNode* nl = new ListNode(i, prev);
        prev = nl;
    }

    std::cout << "\nOld list...\n";
    showList(prev);
    ListNode* hh = ReverseList(prev);
    std::cout << "\n\nReverse list...\n";
    showList(hh);
    ListNode g;
}