from random import randrange
from string import ascii_letters

from hstest import *


class TestUserProgram(StageTest):
    @dynamic_test(data=[("Yeah, okay!", "Great!"), ("Let's be friends.", "What a pity!"), ("No.", None)])
    def test(self, test_answer, user_answer):
        pr = TestedProgram()
        pr.start()

        g, p = randrange(100, 1000), randrange(100, 1000)

        raw_output1 = pr.execute(f"g is {g} and p is {p}")

        self.check_empty_or_none_output(raw_output1)

        if raw_output1.lower().strip() != "ok":
            raise WrongAnswer(
                "Your output should be equal to  \"OK\".\n"
                f"Your output: {raw_output1}.")

        priv_key = self.generate_private_key(p)
        pub_key = self.generate_public_key(g, priv_key, p)

        raw_output2 = pr.execute(f"A is {pub_key}")

        self.check_empty_or_none_output(raw_output2)

        output_lines = raw_output2.strip().split("\n")

        if len(output_lines) == 1:
            raise WrongAnswer(
                "Your program did not print the encrypted message "
                "\"Will you marry me?\".")
        if len(output_lines) > 2:
            raise WrongAnswer(
                "Your program prints extra lines.")

        string_with_user_pub_key, message = output_lines

        try:
            user_pub_key = int(
                string_with_user_pub_key.lower().lstrip("b is ").strip())

        except ValueError:
            raise WrongAnswer(
                "It is not possible to get B from your output.\n"
                f"Your output: {string_with_user_pub_key}.")

        shared_secret = self.compute_shared_secret(user_pub_key, priv_key, p)

        if self.decrypt_message(message.strip(), shared_secret).lower() != "will you marry me?":
            raise WrongAnswer(
                "Unable to decrypt your message.\n"
                f"Your should encrypt \"Will you marry me?\".")

        raw_user_answer = pr.execute(
            self.encrypt_message(test_answer, shared_secret))

        if test_answer == "No.":
            if not raw_user_answer:
                return CheckResult.correct()
            raise WrongAnswer(
                "If the answer is neither \"Yeah, okay!\" nor "
                "\"Let's be friends.\" your program should not print anything, "
                f"but your output is \"{raw_user_answer}\".")

        self.check_empty_or_none_output(raw_user_answer)

        decrypted_user_answer = self.decrypt_message(raw_user_answer.strip(), shared_secret)

        if decrypted_user_answer.lower() != user_answer.lower():
            raise WrongAnswer(
                "Unable to decrypt your answer.\n"
                "Perhaps you made a mistake in the implementation of the protocol.\n"
                f"Your answer should be \"{user_answer}\", "
                f"but your decrypted output is \"{decrypted_user_answer}\".")

        return CheckResult.correct()

    @staticmethod
    def check_empty_or_none_output(output):
        if not output:
            raise WrongAnswer("Your output is empty or None.")

    @staticmethod
    def compute_shared_secret(pub_key, priv_key, p):
        return pow(pub_key, priv_key, p)

    @staticmethod
    def generate_private_key(p):
        return randrange(2, p)

    @staticmethod
    def generate_public_key(g, priv_key, p):
        return pow(g, priv_key, p)

    @staticmethod
    def encrypt_message(msg, key):
        encrypted_msg = []

        for char in msg:
            if char in ascii_letters:
                letter = chr(((ord(char.lower()) - 97) + key) % 26 + 97)
                if char.isupper():
                    letter = letter.upper()
                char = letter

            encrypted_msg.append(char)

        return "".join(encrypted_msg)

    @staticmethod
    def decrypt_message(msg, key):
        decrypted_msg = []

        for char in msg:
            if char in ascii_letters:
                letter = chr((ord(char.lower()) - 97 - key) % 26 + 97)
                if char.isupper():
                    letter = letter.upper()
                char = letter

            decrypted_msg.append(char)

        return "".join(decrypted_msg)


if __name__ == '__main__':
    TestUserProgram().run_tests()
