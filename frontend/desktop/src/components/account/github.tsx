import { Flex, FlexProps, Icon, Tooltip } from '@chakra-ui/react';
import { useQuery } from '@tanstack/react-query';

export default function GithubComponent(props: FlexProps) {
  const { data } = useQuery(
    ['getGithubStar'],
    () => fetch('https://api.github.com/repos/labring/sealos').then((res) => res.json()),
    {
      staleTime: 24 * 60 * 60 * 1000
    }
  );

  return (
    <Flex
      userSelect={'none'}
      w="32px"
      h="32px"
      borderRadius={'50%'}
      justifyContent={'center'}
      alignItems={'center'}
      backgroundColor={'#FFF'}
      cursor={'pointer'}
      fontWeight={500}
      {...props}
      onClick={() => window.open('https://github.com/labring/sealos')}
    >
      <Icon
        xmlns="http://www.w3.org/2000/svg"
        width="20px"
        height="20px"
        viewBox="0 0 20 20"
        fill="white"
      >
        <path d="M10.5149 2.46429C9.50066 2.46429 8.94582 2.57085 8.11017 2.75393C7.6713 2.47429 7.1805 2.2266 6.71141 2.04216C6.16273 1.82643 5.50582 1.64137 4.96522 1.69937C4.73761 1.72379 4.54128 1.83769 4.40747 2.00232C4.34607 2.06873 4.29451 2.14634 4.25617 2.23385C3.78624 3.30652 3.8237 4.37524 4.02876 5.17706C3.70273 5.62136 3.45054 6.03802 3.2796 6.49567C3.06127 7.08016 2.99839 7.66872 2.99839 8.37421C2.99839 10.2018 3.55784 11.6251 4.60368 12.6213C5.34248 13.325 6.27275 13.7665 7.29637 14.0126C7.19262 14.3116 7.14454 14.6131 7.12749 14.9237C7.12244 14.9608 7.11984 14.9987 7.11984 15.0371L7.11984 15.2993C7.08437 15.3069 7.04899 15.3168 7.01389 15.3292C6.29686 15.5828 4.42832 15.4917 3.48217 13.4202C3.29096 13.0015 2.79658 12.8172 2.37795 13.0084C1.95931 13.1996 1.77494 13.6939 1.96615 14.1126C3.19725 16.808 5.63646 17.3073 7.11984 17.0209L7.11984 17.4781C7.11984 17.9383 7.49294 18.3114 7.95318 18.3114C8.41341 18.3114 8.78651 17.9383 8.78651 17.4781L8.7865 15.2177C8.78661 14.8916 8.82066 14.6928 8.87987 14.5341C8.93817 14.3779 9.04207 14.2042 9.26358 13.9629C9.57478 13.6238 9.5522 13.0967 9.21313 12.7854C9.12786 12.7072 9.03069 12.65 8.92805 12.6137C8.85498 12.5841 8.77645 12.5645 8.69396 12.5566C7.41295 12.4343 6.42015 12.0497 5.75319 11.4145C5.10404 10.7961 4.66505 9.84643 4.66505 8.37421C4.66505 7.76687 4.72051 7.40117 4.8409 7.07886C4.96477 6.74723 5.18069 6.39831 5.6054 5.85725C5.68261 5.75888 5.73419 5.64856 5.76109 5.53435C5.79993 5.37587 5.79274 5.20438 5.73033 5.04034C5.57774 4.63929 5.48097 4.05052 5.61339 3.42948C5.75142 3.46674 5.91501 3.5199 6.10155 3.59324C6.57035 3.77757 7.06537 4.04219 7.45271 4.32035C7.58549 4.4157 7.7368 4.46641 7.88875 4.47541C7.9655 4.4805 8.04425 4.47496 8.12315 4.45762L8.25749 4.42806C9.19661 4.22125 9.60665 4.13096 10.5149 4.13096C11.472 4.13096 11.9343 4.22367 12.9808 4.45706C13.0503 4.47256 13.1197 4.47889 13.1878 4.47685C13.3544 4.47581 13.5223 4.42493 13.668 4.32035C14.0553 4.04219 14.5503 3.77757 15.0191 3.59324C15.2035 3.52076 15.3654 3.46799 15.5024 3.4308C15.6309 4.03267 15.5421 4.58336 15.3904 4.98448C15.3739 5.02807 15.3613 5.07218 15.3525 5.11643C15.2813 5.36893 15.3302 5.65098 15.5085 5.86975C15.8603 6.30132 16.0766 6.61787 16.2175 6.96875C16.3557 7.31312 16.4425 7.74564 16.4425 8.42196C16.4425 9.79112 16.0558 10.7082 15.4422 11.328C14.8146 11.962 13.85 12.3861 12.5089 12.5597C12.4949 12.5615 12.4811 12.5637 12.4673 12.5661C12.2703 12.5692 12.0736 12.6417 11.9171 12.7854C11.578 13.0967 11.5554 13.6238 11.8666 13.9629C12.0881 14.2042 12.192 14.3779 12.2503 14.5341C12.3095 14.6928 12.3436 14.8916 12.3437 15.2177L12.3437 17.4781C12.3437 17.9383 12.7168 18.3114 13.177 18.3114C13.6373 18.3114 14.0104 17.9383 14.0104 17.4781L14.0104 15.0371C14.0104 14.9987 14.0078 14.9608 14.0027 14.9237C13.9855 14.6103 13.9367 14.3061 13.831 14.0045C14.9295 13.7266 15.889 13.2458 16.6267 12.5005C17.6174 11.4997 18.1092 10.1226 18.1092 8.42196C18.1092 7.59831 18.0029 6.94238 17.7642 6.34782C17.587 5.90648 17.3488 5.53004 17.0789 5.17087C17.2969 4.35681 17.3282 3.29947 16.8497 2.20714C16.7832 2.05551 16.6771 1.93357 16.5488 1.84818C16.4364 1.76822 16.3024 1.71514 16.1554 1.69937C15.6148 1.64137 14.9579 1.82643 14.4093 2.04216C13.939 2.22706 13.4469 2.47552 13.0072 2.75601C12.1219 2.56301 11.5322 2.46429 10.5149 2.46429Z" />
      </Icon>
    </Flex>
  );
}
