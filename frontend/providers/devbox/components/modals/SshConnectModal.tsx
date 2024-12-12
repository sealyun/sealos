import {
  Box,
  Modal,
  ModalBody,
  ModalContent,
  ModalOverlay,
  ModalHeader,
  ModalCloseButton,
  Flex,
  Text,
  Button,
  Divider,
  Stepper,
  Step,
  StepIndicator,
  StepNumber,
  StepStatus,
  StepSeparator,
  Circle,
  Tabs,
  TabList
} from '@chakra-ui/react'
import { useState } from 'react'
import { useTranslations } from 'next-intl'

import Code from '../Code'
import Tab from '../Tab'
import MyIcon from '../Icon'
import ScriptCode from '../ScriptCode'

interface JetBrainsGuideData {
  devboxName: string
  runtimeType: string
  privateKey: string
  userName: string
  token: string
  workingDir: string
  host: string
  port: string
}

const systemList = ['Windows', 'Mac', 'Linux']

const SshConnectModal = ({
  onClose
}: {
  onSuccess: () => void
  onClose: () => void
  jetbrainsGuideData: JetBrainsGuideData
}) => {
  const t = useTranslations()

  const [activeTab, setActiveTab] = useState(0)

  return (
    <Box>
      <Modal isOpen onClose={onClose} lockFocusAcrossFrames={false}>
        <ModalOverlay />
        <ModalContent top={'5%'} maxWidth={'800px'} w={'700px'} h={'80%'} position={'relative'}>
          <ModalHeader pl={10}>{t('jetbrains_guide_config_ssh')}</ModalHeader>
          <ModalCloseButton top={'10px'} right={'10px'} />
          <ModalBody pb={6} overflowY={'auto'}>
            <Tabs onChange={(index) => setActiveTab(index)} mb={4} colorScheme={'brightBlue'}>
              <TabList>
                {systemList.map((item) => (
                  <Tab key={item}>{item}</Tab>
                ))}
              </TabList>
            </Tabs>
            {/* one-click */}
            <Flex flexDirection={'column'} gap={4}>
              <Text fontSize={'18px'} fontWeight={500} color={'grayModern.900'}>
                {t('jetbrains_guide_one_click_setup')}
              </Text>
              <Text fontSize={'14px'} color={'grayModern.900'} fontWeight={400} lineHeight={'20px'}>
                {t.rich('jetbrains_guide_one_click_setup_desc', {
                  blue: (chunks) => (
                    <Text fontWeight={'bold'} display={'inline-block'} color={'brightBlue.600'}>
                      {chunks}
                    </Text>
                  ),
                  lightColor: (chunks) => (
                    <Text color={'grayModern.600'} display={'inline-block'}>
                      {chunks}
                    </Text>
                  )
                })}
              </Text>
              <Button
                leftIcon={<MyIcon name="download" color={'grayModern.500'} w={'16px'} />}
                w={'fit-content'}
                bg={'white'}
                color={'grayModern.600'}
                border={'1px solid'}
                borderColor={'grayModern.200'}
                borderRadius={'6px'}
                _hover={{
                  color: 'brightBlue.600',
                  '& svg': {
                    color: 'brightBlue.600'
                  }
                }}>
                {t('download_scripts')}
              </Button>
              <ScriptCode />
            </Flex>
            <Divider my={6} />
            {/* step-by-step */}
            <Stepper orientation="vertical" index={-1} mt={4} gap={0} position={'relative'}>
              {/* 1 */}
              <Box w={'100%'}>
                <Step>
                  <StepIndicator backgroundColor={'grayModern.100'} borderColor={'grayModern.100'}>
                    <StepStatus incomplete={<StepNumber />} />
                  </StepIndicator>
                  <Box mt={1} ml={2} mb={5} flex={1}>
                    <Box fontSize={'14px'} mb={3}>
                      {t.rich('jetbrains_guide_download_private_key', {
                        blue: (chunks) => (
                          <Text
                            fontWeight={'bold'}
                            display={'inline-block'}
                            color={'brightBlue.600'}>
                            {chunks}
                          </Text>
                        )
                      })}
                    </Box>
                    <Button
                      leftIcon={<MyIcon name="download" color={'grayModern.600'} w={'16px'} />}
                      bg={'white'}
                      color={'grayModern.600'}
                      borderRadius={'5px'}
                      borderWidth={1}
                      size={'sm'}
                      _hover={{
                        color: 'brightBlue.600',
                        '& svg': {
                          color: 'brightBlue.600'
                        }
                      }}
                      onClick={() => {
                        window.open('https://code-with-me.jetbrains.com/remoteDev', '_blank')
                      }}>
                      {t('download_private_key')}
                    </Button>
                  </Box>
                  <StepSeparator />
                </Step>
              </Box>
              {/* 2 */}
              <Box w={'100%'}>
                <Step>
                  <StepIndicator backgroundColor={'grayModern.100'} borderColor={'grayModern.100'}>
                    <StepStatus incomplete={<StepNumber />} />
                  </StepIndicator>
                  <Flex mt={1} ml={2} mb={5} flex={1} h={'40px'}>
                    <Box fontSize={'14px'}>
                      {t.rich('jetbrains_guide_move_to_path', {
                        blue: (chunks) => (
                          <Text
                            fontWeight={'bold'}
                            display={'inline-block'}
                            color={'brightBlue.600'}>
                            {chunks}
                          </Text>
                        )
                      })}
                    </Box>
                    <Box
                      color={'grayModern.900'}
                      _hover={{
                        color: 'brightBlue.600',
                        '& svg': {
                          color: 'brightBlue.600'
                        }
                      }}
                      cursor={'pointer'}
                      ml={2}>
                      <MyIcon name="copy" color={'grayModern.500'} w={'16px'} />
                    </Box>
                  </Flex>
                  <StepSeparator />
                </Step>
              </Box>
              {/* 3 */}
              <Box w={'100%'}>
                <Step>
                  <StepIndicator backgroundColor={'grayModern.100'} borderColor={'grayModern.100'}>
                    <StepStatus incomplete={<StepNumber />} />
                  </StepIndicator>
                  <Flex mt={1} ml={2} mb={5} flexDirection={'column'} gap={4} flex={1}>
                    <Box fontSize={'14px'}>
                      {t.rich('jetbrains_guide_modified_file', {
                        blue: (chunks) => (
                          <Text
                            fontWeight={'bold'}
                            display={'inline-block'}
                            color={'brightBlue.600'}>
                            {chunks}
                          </Text>
                        )
                      })}
                    </Box>
                    <ScriptCode />
                  </Flex>
                  <StepSeparator />
                </Step>
              </Box>
              {/* 4 */}
              <Box w={'100%'}>
                <Step>
                  <StepIndicator backgroundColor={'grayModern.100'} borderColor={'grayModern.100'}>
                    <StepStatus incomplete={<StepNumber />} />
                  </StepIndicator>
                  <Flex mt={1} ml={2} mb={5} flexDirection={'column'} gap={4} flex={1}>
                    <Box fontSize={'14px'}>{t('jetbrains_guide_command')}</Box>
                    <ScriptCode />
                  </Flex>
                  <StepSeparator />
                </Step>
              </Box>
              {/* done */}
              <Step>
                <Circle size="10px" bg="grayModern.100" top={-3} left={2.5} position={'absolute'} />
              </Step>
            </Stepper>
          </ModalBody>
        </ModalContent>
      </Modal>
    </Box>
  )
}

export default SshConnectModal
