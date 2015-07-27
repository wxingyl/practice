__author__ = 'xing'


class Solution(object):
    def dfs(self, ret_list, s, ret_ip, level):
        if level == 4:
            if len(s) == 0:
                ret_list.append(ret_ip[1:])
            return
        for i in xrange(1, 4):
            if i > len(s):
                continue
            val = s[:i]
            if i != 1 and val[0] == '0':
                continue
            elif i == 3 and val > '255':
                continue
            self.dfs(ret_list, s[i:], ret_ip + '.' + val, level + 1)

    # @param {string} s
    # @return {string[]}
    def restoreIpAddresses(self, s):
        ret = []
        self.dfs(ret, s, '', 0)
        return ret

    def restoreIpAddressesHelper(self, s, n):
        if not n:
            if s:
                number = int(s)
                if number <= 255 and ((number > 0 and s[0] != '0') or len(s) == 1):
                    return [s]
                else:
                    return None
            else:
                return None
        else:
            result = []
            for i in xrange(1, 4):
                if i < len(s):
                    number = int(s[:i])
                    if number <= 255 and ((number > 0 and s[0] != '0') or i == 1):
                        sufix = self.restoreIpAddressesHelper(s[i:], n - 1)
                        if sufix:
                            for item in sufix:
                                result.append(s[:i] + '.' + item)
            return result
