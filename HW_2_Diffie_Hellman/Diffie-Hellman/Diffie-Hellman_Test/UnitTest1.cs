using System.Numerics;
using static Diffie_Hellman.utils.CryptoUtils;

using Xunit;
using Xunit.Abstractions;

namespace Diffie_Hellman_Test
{
    public class UnitTest1
    {
        private readonly ITestOutputHelper _testOutputHelper;

        public UnitTest1(ITestOutputHelper testOutputHelper)
        {
            _testOutputHelper = testOutputHelper;
        }

        [Fact]
        public void ModPowTest()
        {
            var res = ModPow(1123, 232, 12312312);
            Assert.Equal(res, BigInteger.ModPow(12312312, 232, 1123));
        }
    }
}