matasano-challenges
===================
My solutions to the problem posted at http://cryptopals.com/


Where are the solutions ?
^^^^^^^^^^^^^^^^^^^^^^^^^
The code is not organized around the challenges itself. Some are in tests, some are solved through dedicated solvers,
with my primary concern being maximising code re-usage as much as possible. If you want to just see the parts where
the challenge is actually tackled, I marked the points-of-interest by a comment like

.. highlights::
    // Matasano <challenge_number>

You can grep the code to find them.


Why didn't you used golang's built-in library?
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
This repository is also my first code in golang. I thought by neglecting the standard library I would gain much more
insight on how golang works. I discarded everything but some convenience functions and built my own abstractions.



AN IMPORTANT NOTE
^^^^^^^^^^^^^^^^^
This code is NOT cryptographically safe and it will not provide you security you are trying to establish. You are
most welcome inspecting the code for educational purposes and playing with it, but do not, >> DO NOT << use this code
in an environment that is even remotely trying to be secure. Always use well-tested and trusted public libraries for
this purpose.