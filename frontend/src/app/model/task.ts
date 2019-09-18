export interface Task {
    urno: number;
    urno_manager: number;
    urno_company: number;
    urno_project: number;
    urno_service_code: number;
    urno_responsible: number;
    shortinfo: string;
    from_datetime: string;
    until_datetime: string;
    reminder_datetime: string;
    status: string;
    priority: string;
    description: string;
}
