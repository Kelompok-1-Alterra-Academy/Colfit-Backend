package response

import (
	bookingdetails "CalFit/business/booking_details"
	classResponse "CalFit/controllers/classes/response"
	schedulesResponse "CalFit/controllers/schedules/response"
	"time"
)

type Booking_details struct {
	Amount             int                         `json:"amount"`
	Status             string                      `json:"status"`
	UserID             int                         `json:"user_id"`
	OperationalAdminID int                         `json:"operational_admin_id"`
	PaymentID          int                         `json:"payment_id"`
	Class              classResponse.ClassResponse `json:"class"`
	ScheduleID         int                         `json:"schedule_id,omitempty"`
	CreatedAt          *time.Time                  `json:"created_at"`
	UpdatedAt          *time.Time                  `json:"updated_at"`
}

func FromDomain(domain bookingdetails.Domain) Booking_details {
	class := classResponse.ClassResponse{
		ID:   uint(domain.Id),
		Name: domain.ClassName,
		Schedules: []schedulesResponse.Schedules{
			{
				TimeSchedule: domain.TimeSchedule,
			},
		},
	}
	return Booking_details{
		Amount:             domain.Amount,
		Status:             domain.Status,
		UserID:             domain.UserID,
		OperationalAdminID: domain.OperationalAdminID,
		PaymentID:          domain.PaymentID,
		Class:              class,
		ScheduleID:         domain.ScheduleID,
		CreatedAt:          &domain.CreatedAt,
		UpdatedAt:          &domain.UpdatedAt,
	}
}
