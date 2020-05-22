## 快速幂算法

   999的二进制为1111100111,3^999 等于3^(512+256+128+644+63+4+2+1)，刚好与各二进制位对应，于是有求幂算法：

   ```c++
        double power(double x, int n) {
                double result = 1;
                while(n) {
                        if(n&1)
                                result *= x;
                        x *= x;
                        n >>= 1;

                }
                return result;
        }


   ```

## 递归快速求幂

```c++
        double power(double x,  int n) {
                bool isNag = false;
                if (n == 0)
                        return 1;
                if (n == 1)
                        return base;
                if (n < 0 )
                        isNag = true;
                unsigned int exp = n > 0? n:-n;
                double result = abs_power(x,exp);
                if (isNag)
                        return 1/result;
                else 
                                return result;
        }
        double abs_power(double base, unsigned int n) {
                if (n & 1)
                        return abs_power(base*base, n >> 1)*base;
                else
                        return abs_power(base*base,n >> 1);

        }

```