


.. _scripts/modtools:

========================
modtools/* - for modders
========================
Scripts which provide tools for modders, often with changes to the raw files.
Not intended to be called manually by end-users.

These scripts are mostly useful for raw modders and scripters.

They all have
standard arguments: arguments are of the form ``tool -argName1 argVal1
-argName2 argVal2``. This is equivalent to ``tool -argName2 argVal2 -argName1
argVal1``. It is not necessary to provide a value to an argument name: ``tool
-argName3`` is fine. Supplying the same argument name multiple times will
result in an error. Argument names are preceded with a dash. The ``-help``
argument will print a descriptive usage string describing the nature of the
arguments. For multiple word argument values, brackets must be used: ``tool
-argName4 [ sadf1 sadf2 sadf3 ]``. In order to allow passing literal braces as
part of the argument, backslashes are used: ``tool -argName4 [ \] asdf \foo ]``
sets ``argName4`` to ``] asdf foo``. The ``*-trigger`` scripts have a similar
policy with backslashes.

.. include:: include-all.rst

.. toctree::
   :hidden:
   
   include-all

