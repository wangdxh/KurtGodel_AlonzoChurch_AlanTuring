### 哥德尔 邱奇 图灵 

哥德尔的伟大之处在于：  
1.哥德尔数是最早的解释器的原型；  
2.哥德尔的对角线引理产生了自引用，虽然哥德尔声称在康托尔的对角线中找到了灵感，但是基本上看不出它们之间的关系。
  
数学中，你最服的技巧是哪个？  
知乎的一个回答:  
https://www.zhihu.com/question/66094972/answer/238396535  


数理逻辑虽然渐渐掉出了主流数学系的关注, 但因为哥德尔不完备定理实在是太出名了, 而且其中哥德尔达到self-reference的引理也很少被人提及, 所以这里写一下证明不完备定理中必不可少的一个引理: Gödel's Diagonal Lemma (a.k.a. Gödel's Fixed Point Theorem)

引理内容：在一个足够强(suffciently strong, 即可以做到Gödel numbering)的算术体系里, 对于每一个formula F(x), 都存在一个formula A, 使得A逻辑等价于F(|A|).

这里|A|表示的是formula A的Gödel number. 非形式上来讲, A所表达的就是“我是F的”.

背景：这个引理是在介绍完Gödel numbering之后引入的. Gödel numbering的作用是使算术体系中的句法对象(加号, 等于号, 一段证明等)都能被赋予独一无二的自然数, 也就相当于让算术体系能用自己的语言描述出自己的句法内容. 与本文相关的句法对象是“Proof(x, y)”, 即“x is the Gödel number of a proof of y", 在这之上我们可以再定义一个provability predicate, Prov(x), 即“there exists some number n such that n is the Gödel number of a proof of x". 注意这里只是让算术体系能够有办法描述自己体系中的句法对象而已, 还没有达到self-reference的地步. 在哥德尔之前, 数学界的共识基本上是self-reference是自然语言不严谨的结果, 严谨的数学语言不会出现. 哥德尔的引理则否定了这种认识：即在足够强的算术体系中, self-reference无法避免.

证明： 对于任意formula F(x), 皆存在formula A, 使得A<->F(|A|). （“|x|”指代的是x的Gödel number. ）

1. 当算术体系T足够强时, T将可以构造一个“apply(x, y)” function, 使得对于任意formula A(x)与B(x), apply(|A(x)|, |B(x)|)=|A(|B(x)|)|
2. 给定任意formula F(x), 定义G(x)为F(apply(x, x)).
3. 定义A为G(|G(x)|).
4. A <-> G(|G(x)|)
<-> F(apply(|G(x)|, |G(x)|)
<-> F(|G(|G(x)|)|)
<-> F(|A|)
证毕.

哥德尔不完备定理则是通过该引理构造出A<->not-Prov(|A|)来证明的. 这里的A非形式上表达的就是“我是无法证明的”).

这个trick在接下来几十年内出现的各种限制性定理证明中都起到了不可或缺的作用. 例如Tarski's Indefinability Theorem ("a truth predicate of a sufficiently strong system cannot be defined in that system")的证明也是同样利用该引理构造出A<->not-True(|A|). A在这里非形式上表达的就是“我是假的”.
