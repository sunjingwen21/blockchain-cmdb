import React from 'react';
import { Layout, Typography, Card, Statistic, Row, Col } from 'antd';
import { DatabaseOutlined, BlockOutlined, SafetyOutlined } from '@ant-design/icons';

const { Header, Content, Footer } = Layout;
const { Title } = Typography;

const Dashboard: React.FC = () => {
  return (
    <Layout style={{ minHeight: '100vh' }}>
      <Header style={{ background: '#fff', padding: '0 24px' }}>
        <Title level={3} style={{ margin: '16px 0' }}>
          <DatabaseOutlined /> Blockchain CMDB
        </Title>
      </Header>
      
      <Content style={{ padding: '24px', background: '#f0f2f5' }}>
        <Row gutter={[16, 16]}>
          <Col span={8}>
            <Card>
              <Statistic
                title="Total Assets"
                value={0}
                prefix={<DatabaseOutlined />}
              />
            </Card>
          </Col>
          <Col span={8}>
            <Card>
              <Statistic
                title="Blockchain Records"
                value={0}
                prefix={<BlockOutlined />}
              />
            </Card>
          </Col>
          <Col span={8}>
            <Card>
              <Statistic
                title="Security Status"
                value="Secure"
                prefix={<SafetyOutlined />}
                valueStyle={{ color: '#3f8600' }}
              />
            </Card>
          </Col>
        </Row>

        <Card style={{ marginTop: 24 }}>
          <Title level={4}>Welcome to Blockchain CMDB</Title>
          <p>This is a production-grade Configuration Management Database platform with blockchain integration.</p>
          <p>Features coming soon:</p>
          <ul>
            <li>Asset management with blockchain audit trail</li>
            <li>Real-time dashboard and analytics</li>
            <li>Multi-signature approval workflows</li>
            <li>Immutable configuration history</li>
          </ul>
        </Card>
      </Content>

      <Footer style={{ textAlign: 'center' }}>
        Blockchain CMDB Platform ©2026 Created by OpenClaw Agent
      </Footer>
    </Layout>
  );
};

export default Dashboard;
