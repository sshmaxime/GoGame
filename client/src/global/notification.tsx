import { notification } from 'antd';

class NotificationCenter {
  private static instance: NotificationCenter;

  private constructor() { }
  public static getInstance(): NotificationCenter {
    if (!NotificationCenter.instance) {
      NotificationCenter.instance = new NotificationCenter();
    }
    return NotificationCenter.instance;
  }

  notificationErr = (error: string) => {
    notification["error"]({
      message: "Error",
      description: error,
    });
  }
  notificationSuccess = (data: string) => {
    notification["success"]({
      message: "Success !",
      description: data,
    });
  }
}

export default NotificationCenter;




