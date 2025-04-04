import os
import requests

cookies = {
    'passport_csrf_token': '3a218467d059ba2d72eb5c54a99638d4',
    'passport_csrf_token_default': '3a218467d059ba2d72eb5c54a99638d4',
    'ticket_guard_has_set_public_key': '1',
    'd_ticket': 'df117dbc38bcb4fbbd89cf21b62669670be58',
    'multi_sids': '1115358952428345%3A2c2476fd01f18f6c9a5c7615fa857072',
    'n_mh': 'EJpGAl8UyrYuf8qp_T0CjxfWrnBWnkqZm_Fq-wJea5U',
    'odin_tt': '21ea15607196e79fa110c11b80466f643499f1c90d7f70cfc28c7c056e7845df02525a7a1e92fdbf54a38ed92a8ca0725ebfdbcf5c3097232b44727a3f79d5b3efb345189880db01fd8fcbe5a12b8496',
    'passport_assist_user': 'CkELlcduQ6VcQe7ureljsdz4XrBrFT6Ggue3kqykxn_x5Jk9uRWr98Ntxl6aeHuzsYvkhIcfKtjgLmt6zIBiiPoWkxpKCjx5_Er76C7Pa4u29aURxtxTTbX1SApjMsmQy5OmD_mGkLCh-h5KZw0ohNJqK32SMYaRA7M2M0Qjmy-km1cQzMPXDRiJr9ZUIAEiAQN7Miu_',
    'sessionid': '2c2476fd01f18f6c9a5c7615fa857072',
    'sessionid_ss': '2c2476fd01f18f6c9a5c7615fa857072',
    'sid_guard': '2c2476fd01f18f6c9a5c7615fa857072%7C1721825582%7C5184000%7CSun%2C+22-Sep-2024+12%3A53%3A02+GMT',
    'sid_tt': '2c2476fd01f18f6c9a5c7615fa857072',
    'uid_tt': '7ee5995032326a3704ca3789e090f936',
    'uid_tt_ss': '7ee5995032326a3704ca3789e090f936',
    'install_id': '1871824318410835',
    'ttreq': '1' + os.getenv('275454359d83880e4765589dfefebb1041cc3c84', ''),
    'store-region': 'cn-fj',
    'store-region-src': 'uid',
}

headers = {
    'Host': 'api5-normal-lq.amemv.com',
    'Connection': 'keep-alive',
     'Content-Length': '99',
    'x-tt-multi-sids': '1115358952428345%3A2c2476fd01f18f6c9a5c7615fa857072',
    'X-SS-Cookie': 'store-region=cn-fj; store-region-src=uid; install_id=1871824318410835; ttreq=1' + os.getenv('275454359d83880e4765589dfefebb1041cc3c84', '') + '; multi_sids=1115358952428345%3A2c2476fd01f18f6c9a5c7615fa857072; n_mh=EJpGAl8UyrYuf8qp_T0CjxfWrnBWnkqZm_Fq-wJea5U; odin_tt=21ea15607196e79fa110c11b80466f643499f1c90d7f70cfc28c7c056e7845df02525a7a1e92fdbf54a38ed92a8ca0725ebfdbcf5c3097232b44727a3f79d5b3efb345189880db01fd8fcbe5a12b8496; passport_assist_user=CkELlcduQ6VcQe7ureljsdz4XrBrFT6Ggue3kqykxn_x5Jk9uRWr98Ntxl6aeHuzsYvkhIcfKtjgLmt6zIBiiPoWkxpKCjx5_Er76C7Pa4u29aURxtxTTbX1SApjMsmQy5OmD_mGkLCh-h5KZw0ohNJqK32SMYaRA7M2M0Qjmy-km1cQzMPXDRiJr9ZUIAEiAQN7Miu_; sessionid=2c2476fd01f18f6c9a5c7615fa857072; sessionid_ss=2c2476fd01f18f6c9a5c7615fa857072; sid_guard=2c2476fd01f18f6c9a5c7615fa857072%7C1721825582%7C5184000%7CSun%2C+22-Sep-2024+12%3A53%3A02+GMT; sid_tt=2c2476fd01f18f6c9a5c7615fa857072; uid_tt=7ee5995032326a3704ca3789e090f936; uid_tt_ss=7ee5995032326a3704ca3789e090f936; d_ticket=df117dbc38bcb4fbbd89cf21b62669670be58; ticket_guard_has_set_public_key=1; passport_csrf_token=3a218467d059ba2d72eb5c54a99638d4; passport_csrf_token_default=3a218467d059ba2d72eb5c54a99638d4',
    'sdk-version': '2',
    'Content-Type': 'application/x-www-form-urlencoded',
    'x-Tt-Token': '002c2476fd01f18f6c9a5c7615fa85707200aa4004bd57a564f12e66e9118999b282668eec4eb2264190e52a79fa6063974cd7b58a5a6a7958a03a01de9ab48b131a6d00421272452805d894ba18fcedcc519fc8db6872420818e0213971318301f94-1.0.1',
    'User-Agent': 'Aweme 19.0.0 rv:190015 (iPhone; iOS 13.5.1; zh_CN) Cronet',
    'x-vc-bdturing-sdk-version': '2.2.3',
    'tt-request-time': '1721825902361',
     'Cookie': 'passport_csrf_token=3a218467d059ba2d72eb5c54a99638d4; passport_csrf_token_default=3a218467d059ba2d72eb5c54a99638d4; ticket_guard_has_set_public_key=1; d_ticket=df117dbc38bcb4fbbd89cf21b62669670be58; multi_sids=1115358952428345%3A2c2476fd01f18f6c9a5c7615fa857072; n_mh=EJpGAl8UyrYuf8qp_T0CjxfWrnBWnkqZm_Fq-wJea5U; odin_tt=21ea15607196e79fa110c11b80466f643499f1c90d7f70cfc28c7c056e7845df02525a7a1e92fdbf54a38ed92a8ca0725ebfdbcf5c3097232b44727a3f79d5b3efb345189880db01fd8fcbe5a12b8496; passport_assist_user=CkELlcduQ6VcQe7ureljsdz4XrBrFT6Ggue3kqykxn_x5Jk9uRWr98Ntxl6aeHuzsYvkhIcfKtjgLmt6zIBiiPoWkxpKCjx5_Er76C7Pa4u29aURxtxTTbX1SApjMsmQy5OmD_mGkLCh-h5KZw0ohNJqK32SMYaRA7M2M0Qjmy-km1cQzMPXDRiJr9ZUIAEiAQN7Miu_; sessionid=2c2476fd01f18f6c9a5c7615fa857072; sessionid_ss=2c2476fd01f18f6c9a5c7615fa857072; sid_guard=2c2476fd01f18f6c9a5c7615fa857072%7C1721825582%7C5184000%7CSun%2C+22-Sep-2024+12%3A53%3A02+GMT; sid_tt=2c2476fd01f18f6c9a5c7615fa857072; uid_tt=7ee5995032326a3704ca3789e090f936; uid_tt_ss=7ee5995032326a3704ca3789e090f936; install_id=1871824318410835; ttreq=1' + os.getenv('275454359d83880e4765589dfefebb1041cc3c84', '') + '; store-region=cn-fj; store-region-src=uid',
    'x-tt-passport-csrf-token': '3a218467d059ba2d72eb5c54a99638d4',
    'x-tt-dt': 'AAAQ4WJS73ZVVSNCWESBA54VFQB4RCHYFLHJ3DD72RDGXDZKKQ43B4ZFGNFKUMNPZSBW42CPGD52JKJX4SYK346NIPPZOHCGAYEJBN325ULZY3XUXQUSXUAKGDSVS',
    'passport-sdk-version': '5.12.1',
    'X-SS-STUB': '28DF655EB957A67C17694DB9C36EFD79',
    'x-bd-kmsv': '1',
    'x-tt-request-tag': 's=1',
    'X-SS-DP': '1128',
    'x-tt-trace-id': '00-e4d240b30dde665daec2b8b01fb30468-e4d240b30dde665d-01',
    # 'Accept-Encoding': 'gzip, deflate, br',
    'X-Argus': 'Mc+UNgplTZwpTTQ6zx/VWmzMcs483CHkHalpP8JkZiB+x8erMjTBxgjUFvKYWxJw9C3DCsJ1QUfs+/IXnU9SfQxAtgXpRRfNw2BFyFIx6K0klpn+IdGMZLDAWgwawhvnHlTdDcmD+gY/Ma7eIhBLeu2Mv0P/43vvLECu19PLC79HSRnfjEujAhDBI2fYPmBbleLVgVCaWkFOFWisTR+MbtnQOlE/DQeCbaxe/lF8thq8JQh9kBGWusO5SGh6eZ+WM2TNEizyb0LAtBOKobPYuSwqQvttMooZSGIYvME8wNzJSxwXiDrBZmIeseDHyvwNBO3ipowtCswesmS7qzk3RQdwt5y66FP6QnGVWmTIf06KlTWxNQyBnHvx7Os1fScSr4wsDIUtlzWhE5VPfqxa3ShusJXNgQv4kb4Qv/b+Z7/wAipiJ0SoWQO81I1zhyuspkmYttMqiBBC/4hEHqE8znvDerGn4UbTG6ozXNQBSH5uYtUcSQ37hRU9hiAFFUYYCDq5bZ05D6qh9SIwB7UMo/PrdWsw9LDfQXDK9grV0qCJiA==',
    'X-Gorgon': '840460622001afdbe42ea1a0701ffaccbd5216c2c7fd39675cc1',
    'X-Khronos': '1721825902',
    'X-Ladon': 'EPMeIFJKyE0rJ5xHpMAuMBkkFggrzlgLVxyjzj0Pnxvxz5qA',
}

params = {
    'device_id': '3912499836234635',
    'os_version': '13.5.1',
    'multi_login': '1',
    'app_id': '1128',
    'iid': '1871824318410835',
    'app_name': 'aweme',
    'slide_guide_has_shown': '1',
    'appTheme': 'light',
    'ac': 'WIFI',
    'js_sdk_version': '2.36.1.0',
    'ssmix': 'a',
    'version_code': '19.0.0',
    'channel': 'App Store',
    'is_vcd': '1',
    'tma_jssdk_version': '2.36.1.0',
    'os_api': '18',
    'need_personal_recommend': '1',
    'install_id': '1871824318410835',
    'device_platform': 'iphone',
    'device_type': 'iPhone9,3',
    'build_number': '190015',
    'is_guest_mode': '0',
    'minor_status': '0',
    'package': 'com.ss.iphone.ugc.Aweme',
    'mcc_mnc': '',
    'aid': '1128',
    'screen_width': '750',
    'cdid': 'C155A3ED-4936-4ADF-8234-16740EF0E818',
    'app_version': '19.0.0',
    'resolution': '750*1334',
}

data = {
    'is_vcd': '1',
    'mix_mode': '1',
    'mobile': '2e3d333430353235313630333535',
    'multi_login': '1',
    'password': '74323736313d3c323137',
}

response = requests.post(
    'https://api5-normal-lq.amemv.com/passport/mobile/login/',
    params=params,
    cookies=cookies,
    headers=headers,
    data=data,
    verify=False,
)
print(response.text())