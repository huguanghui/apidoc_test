/** @brief   基础头文件
 *  @file    demo.h
 *  @author  huguanghui
 *  @version 1.0.0
 *  @date    2019-04-09
 *  @note    公用接口
 *  @since   基础功能
 */

enum
{
    OST_PLATFORM_WIN32         = 1,
    OST_PLATFORM_LINUX_X86     = 2,
    OST_PLATFORM_LINUX_ARM     = 3,
    OST_PLATFORM_ANDROID       = 4,
    OST_PLATFORM_MACOSX        = 5,
};

/**
* @brief 类的简单概述 \n
* 类的详细概述
*/
class Example
{
    int a;
};

/** 
* @brief 函数简要说明-测试函数
* @param index    参数1
* @param t        参数2 @see CTest
*
* @return 返回说明
*        -<em>false</em> fail
*        -<em>true</em> succeed
*/
bool Test(int index, const CTest& t);