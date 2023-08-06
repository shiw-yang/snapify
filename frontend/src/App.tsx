import { useEffect, useState } from 'react';
// import logo from './assets/images/logo-universal.png';
import './App.css';
import { GetMenusName } from "../wailsjs/go/menu/MenuModule";
import MenuItem from "antd/lib/menu/MenuItem";
import { Layout, Space, theme, Menu } from 'antd';
const { Header, Content } = Layout;

const App: React.FC = () => {
    const [menus, setMenus] = useState<string[]>([])
    useEffect(() => {
        GetMenusName().then((res) => {
            setMenus(res)
        })
    }, [])
    const {
        token: { colorBgContainer }
    } = theme.useToken();
    return (
        <Layout className="layout">
            <Header className='header'>
                <Menu
                    style={{ "margin-left": "auto" }}
                    theme="dark"
                    mode="horizontal"
                    defaultSelectedKeys={['2']}
                    items={new Array(menus.length).fill(menus).map(
                        (item, index) => {
                            return {
                                key: index,
                                label: item[index],
                            }
                        }
                    )}
                />
            </Header>
        </Layout >
    );
}
export default App
